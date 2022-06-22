package services_test

import (
	"fmt"
	"github.com/gitcpu-io/origin/models/weather"
	"github.com/gitcpu-io/origin/services/mocks"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Weather", func() {
	It("first test", func() {
		ctrl := gomock.NewController(GinkgoT())
		client := mocks.NewMockWeatherer(ctrl)
		var result []*weather.Weather
		client.EXPECT().List(gomock.Any(),gomock.Any()).Return(result,nil)
		fmt.Println("success")
	})
})
