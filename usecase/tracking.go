package usecase

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Firmwave11/otten-coffee/models"
	"github.com/PuerkitoBio/goquery"
)

type response struct {
	Status struct {
		Code    string `json:"code"`
		Message string `json:"error_message"`
	} `json:"status"`
	Data interface{} `json:"data"`
}

var (
	url = "https://gist.githubusercontent.com/nubors/eecf5b8dc838d4e6cc9de9f7b5db236f/raw/d34e1823906d3ab36ccc2e687fcafedf3eacfac9/jne-awb.html"
)

const dateLayout = "2006-01-02T15:04:05+07:00"
const dateLayout2 = "02-01-2006 15:04"
const datelayout3 = "02 Februari 2006, 15:04 WIB"

func (r *uc) Tracking(ctx context.Context) (context.Context, interface{}, int, error) {
	var (
		errRes response
	)
	dataTracking := models.Data{}
	errRes.Status.Code = "060101"
	dataTracking.Status.Code = "060101"
	dataTracking.Status.Message = "Delivery tracking detail fetched successfully"
	/* ctx, res, code, err := request.Curl(ctx, url, http.MethodGet, 10*time.Second, nil, nil) */

	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		errRes.Status.Message = fmt.Sprint(err)
		log.Print(err)
		return ctx, errRes, http.StatusBadRequest, err

	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		errRes.Status.Message = "Delivery tracking detail fetched failed"
		log.Printf("status code error: %d %s", res.StatusCode, res.Status)
		return ctx, errRes, res.StatusCode, err
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		dataTracking.Status.Message = fmt.Sprint(err)
		log.Print(err)
		return ctx, dataTracking, http.StatusUnsupportedMediaType, err
	}

	// Find the review items
	doc.Find(".table_style").Each(func(i int, s *goquery.Selection) {
		title := s.Find("thead tr")
		body := s.Find("tbody tr")
		title.Find("td").Each(func(ix int, td *goquery.Selection) {
			if td.Text() == "History " {
				result := models.Historeis{}
				body.Find("td").Each(func(jx int, tdBody *goquery.Selection) {
					fmt.Println(tdBody.Text(), jx)
					if jx%2 != 0 {
						result.Description = tdBody.Text()
						dataTracking.Data.Histories = append(dataTracking.Data.Histories, result)
					} else {
						date, _ := time.Parse(dateLayout2, tdBody.Text())
						result.Createdat = date.Format(dateLayout)
						result.Formatted.Createdat = date.Format(datelayout3)
					}
				})
			}
		})
	})

	for i, j := 0, len(dataTracking.Data.Histories)-1; i < j; i, j = i+1, j-1 {
		dataTracking.Data.Histories[i], dataTracking.Data.Histories[j] = dataTracking.Data.Histories[j], dataTracking.Data.Histories[i]
	}

	lastArr := dataTracking.Data.Histories[0].Description
	splitString := strings.Split(lastArr, " ")

	if splitString[0] == "DELIVERED" {
		dataTracking.Data.Receivedby = strings.Replace(splitString[2], "[", "", -1) + " " + splitString[3]
	}

	fmt.Println(dataTracking.Data.Histories)
	return ctx, dataTracking, http.StatusOK, nil
}
