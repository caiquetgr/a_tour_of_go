package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	data := Input{
		A: []int{10, 20, 30},
		B: []int{50, 10},
	}

	finalResult, err := GatherAndProcess(context.Background(), data)
	if err != nil {
		fmt.Println("aaaaaaa", err)
	} else {
		fmt.Println(finalResult)
	}
}

type AOut int
type BOut int
type COut int
type CIn []int

type Input struct {
	A []int
	B []int
}

type processor struct {
	outA chan AOut
	outB chan BOut
	outC chan COut
	inC  chan CIn
	errs chan error
}

func GatherAndProcess(ctx context.Context, data Input) (COut, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
	defer cancel()

	p := processor{
		outA: make(chan AOut, 1),
		outB: make(chan BOut, 1),
		outC: make(chan COut, 1),
		inC:  make(chan CIn, 1),
		errs: make(chan error, 2),
	}

	p.launch(ctx, data)

	inputC, err := p.waitForAB(ctx)
	if err != nil {
		return 0, err
	}
	p.inC <- inputC

	out, err := p.waitForC(ctx)
	return out, err
}

func (p processor) launch(ctx context.Context, data Input) {

	go func() {
		AResult, err := process(data.A)
		if err != nil {
			p.errs <- err
			return
		}
		p.outA <- AOut(AResult)
	}()

	go func() {
		BResult, err := process(data.B)
		if err != nil {
			p.errs <- err
			return
		}
		p.outB <- BOut(BResult)
	}()

	go func() {
		select {
		case <-ctx.Done():
			return
		case inputC := <-p.inC:
			CResult, err := process(inputC)
			if err != nil {
				p.errs <- err
				return
			}
			p.outC <- COut(CResult)
		}
	}()
}

func (p processor) waitForAB(ctx context.Context) (CIn, error) {

	cIn := CIn{}
	count := 0

	for count < 2 {
		select {
		case a := <-p.outA:
			count++
			cIn = append(cIn, int(a))
		case b := <-p.outB:
			count++
			cIn = append(cIn, int(b))
		case err := <-p.errs:
			return CIn{}, err
		case <-ctx.Done():
			return CIn{}, ctx.Err()
		}
	}

	return cIn, nil
}

func (p processor) waitForC(ctx context.Context) (COut, error) {
	select {
	case cOut := <-p.outC:
		return cOut, nil
	case <-ctx.Done():
		return 0, ctx.Err()
	case err := <-p.errs:
		return 0, err
	}
}

func process(values []int) (int, error) {
	//time.Sleep(40 * time.Millisecond)
	var result int
	for _, value := range values {
		result += value
	}
	return result, nil
}
