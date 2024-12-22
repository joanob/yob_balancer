package main

import (
	"fmt"
	"testing"
)

func Test_GetResourceCombinations(t *testing.T) {
	testCases := []struct {
		desc               string
		resource           ProductionResourceParameters
		resourceParameters []ResourceParameter
		expectedLength     int
	}{
		{
			desc: "Resource with fixed price",
			resource: ProductionResourceParameters{
				Id:  "electricity",
				Min: 1,
				Max: 5,
			},
			resourceParameters: []ResourceParameter{
				{
					Id:  "electricity",
					Min: 1,
					Max: 1,
				},
			},
			expectedLength: 5,
		},
		{
			desc: "Resource with variable price",
			resource: ProductionResourceParameters{
				Id:  "electricity",
				Min: 1,
				Max: 5,
			},
			resourceParameters: []ResourceParameter{
				{
					Id:  "electricity",
					Min: 1,
					Max: 5,
				},
			},
			expectedLength: 25,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			resourceParameters = tC.resourceParameters

			combinations := getResourceCombinations(tC.resource)

			fmt.Println(combinations)

			if len(combinations) != tC.expectedLength {
				t.Error(combinations)
			}
		})
	}
}

func Test_GetResourcesCombinations(t *testing.T) {
	testCases := []struct {
		desc               string
		resources          []ProductionResourceParameters
		resourceParameters []ResourceParameter
		expectedLength     int
	}{
		{
			desc: "One resource",
			resources: []ProductionResourceParameters{
				{
					Id:  "electricity",
					Min: 1,
					Max: 5,
				},
			},
			resourceParameters: []ResourceParameter{
				{
					Id:  "electricity",
					Min: 1,
					Max: 1,
				},
			},
			expectedLength: 5,
		},
		{
			desc: "Two resources",
			resources: []ProductionResourceParameters{
				{
					Id:  "electricity",
					Min: 1,
					Max: 2,
				}, {
					Id:  "water",
					Min: 1,
					Max: 2,
				},
			},
			resourceParameters: []ResourceParameter{
				{
					Id:  "electricity",
					Min: 1,
					Max: 1,
				},
				{
					Id:  "water",
					Min: 2,
					Max: 5,
				},
			},
			expectedLength: 16,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			resourceParameters = tC.resourceParameters

			combinations := make([][]ProcessResource, 0)

			combinations = getResourcesCombinations(combinations, tC.resources)

			fmt.Println(combinations)

			if len(combinations) != tC.expectedLength {
				t.Error(len(combinations))
			}
		})
	}
}
