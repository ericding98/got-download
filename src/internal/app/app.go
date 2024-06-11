package app

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

var (
	season1Episodes []string = []string{
		"https://api.nivodz.com/play/mud.m3u8/WEB/1.0?ce=01c1a12d2a2a6542045919fa6cc161ebefed15c6e79a44b4eb43824c1122138db8c6d85267293bcea43f28b551a961be741de73af837ab5fcb33ee6c7fa74c24d9cfad3f6dec2d7c7bcd5d1f5ea751003ab11a7b529ef64865f5388d7fe43b416245059b217b32a9c00f817ad7f2e6d9358092e6438bc892&pf=3&uk=10a7a3bd516c623e8ff8ff6231956fed&rx=11230&expire=1676685081273&ip=142.188.84.231&sign=e73f698821e6c03d09e55e59a534cdd1&_ts=1676659881273",
		"https://api.nivodz.com/play/mud.m3u8/WEB/1.0?ce=01c1a12d2a2a6542045919fa6cc161ebefed15c6e79a44b4eb43824c1122138d9c16212605c41149f16996117891976c189cadd9b3e76e71d29ac41124548dc23f0718d8b14c3c667bcd5d1f5ea751003ab11a7b529ef64865f5388d7fe43b416245059b217b32a9c00f817ad7f2e6d9358092e6438bc892&pf=3&uk=550bdedb69488832a78c984c78d69153&rx=15430&expire=1676686411104&ip=142.188.84.231&sign=c3514157ff42bc53e58f1f903fe9106d&_ts=1676661211104",
		"https://api.nivodz.com/play/mud.m3u8/WEB/1.0?ce=01c1a12d2a2a6542045919fa6cc161ebefed15c6e79a44b4eb43824c1122138df0aec59e87522da82b9b6dcdc1a037ed9df4bdf9cf6b62060d423cec82f1c0b7007da324a99b9bbf7bcd5d1f5ea751003ab11a7b529ef64865f5388d7fe43b416245059b217b32a9c00f817ad7f2e6d9358092e6438bc892&pf=3&uk=faf78d2297125633bd4feed93b9c9556&rx=846&expire=1676686449101&ip=142.188.84.231&sign=90b8dcddea3208e60e87c4ff54a78fdb&_ts=1676661249101",
		"https://api.nivodz.com/play/mud.m3u8/WEB/1.0?ce=01c1a12d2a2a6542045919fa6cc161ebefed15c6e79a44b4eb43824c1122138d92812622056a8de68d31cfd08adc766b981d55eba8cdd0dc816c9e7d126eac6e955440fa6ced884f7bcd5d1f5ea751003ab11a7b529ef64865f5388d7fe43b416245059b217b32a9c00f817ad7f2e6d9358092e6438bc892&pf=3&uk=fb5386293e50784dc2b64275387e18c1&rx=11970&expire=1676686478774&ip=142.188.84.231&sign=dfad2fde3b862e9bb4c312c77dab467e&_ts=1676661278774",
		"https://api.nivodz.com/play/mud.m3u8/WEB/1.0?ce=01c1a12d2a2a6542045919fa6cc161ebefed15c6e79a44b4eb43824c1122138d05a0d891fa6dcb682c21397012f7f0023ebc9d6e2e9f7944afe0c2c071587db28c21a58ce22ed0af7bcd5d1f5ea751003ab11a7b529ef64865f5388d7fe43b416245059b217b32a9c00f817ad7f2e6d9358092e6438bc892&pf=3&uk=76d9b8afb5185dc281cb27d8906fb5cd&rx=5968&expire=1676686508448&ip=142.188.84.231&sign=4347fd66283873b570db018f65da7c06&_ts=1676661308448",
		"https://api.nivodz.com/play/mud.m3u8/WEB/1.0?ce=01c1a12d2a2a6542045919fa6cc161ebefed15c6e79a44b4eb43824c1122138d3678f944888609b8643f796a1c2f4b3cec1d1ba147aca6ed0b17907adbe1b032245a4ac4f434ce217bcd5d1f5ea751003ab11a7b529ef64865f5388d7fe43b416245059b217b32a9c00f817ad7f2e6d9358092e6438bc892&pf=3&uk=d8c1710243c36db45bf2e3444eaa60d8&rx=19872&expire=1676686587619&ip=142.188.84.231&sign=e418443183c58bbf92cfea13afa3b517&_ts=1676661387619",
		"https://api.nivodz.com/play/mud.m3u8/WEB/1.0?ce=01c1a12d2a2a6542045919fa6cc161ebefed15c6e79a44b4eb43824c1122138d97e13310176af099f15bcb8fb2c1111b72182d752b57d1e7ef061053ba9e10b0222d831252f2ef027bcd5d1f5ea751003ab11a7b529ef64865f5388d7fe43b416245059b217b32a9c00f817ad7f2e6d9358092e6438bc892&pf=3&uk=54c44b91156d78f219b5d1d79fd308ee&rx=4594&expire=1676686622729&ip=142.188.84.231&sign=1354cbf119087075d2a25b18c04c7d33&_ts=1676661422729",
		"https://api.nivodz.com/play/mud.m3u8/WEB/1.0?ce=01c1a12d2a2a6542045919fa6cc161ebefed15c6e79a44b4eb43824c1122138defe84d7a93b4fd5f6ba315a594ac56a132d1dbe1bb570d62da6370294df22a787e6c40402dbcd0217bcd5d1f5ea751003ab11a7b529ef64865f5388d7fe43b416245059b217b32a9c00f817ad7f2e6d9358092e6438bc892&pf=3&uk=344b38794a873ad8631256a6c70e2cb3&rx=2306&expire=1676686655199&ip=142.188.84.231&sign=67f13666c74952f5dfb8452cd001263a&_ts=1676661455199",
		"https://api.nivodz.com/play/mud.m3u8/WEB/1.0?ce=01c1a12d2a2a6542045919fa6cc161ebefed15c6e79a44b4eb43824c1122138df09a236b27a3b759699bbd3ffe6152b1c5571bb9d0e4c371e7ce08a2bbc309dc9baf2a4782abbbc87bcd5d1f5ea751003ab11a7b529ef64865f5388d7fe43b416245059b217b32a9c00f817ad7f2e6d9358092e6438bc892&pf=3&uk=ae068948e38daa616d331377134eafda&rx=9264&expire=1676686690680&ip=142.188.84.231&sign=f761305ee06d0b7984bff3b1e1d2200e&_ts=1676661490680",
		"https://api.nivodz.com/play/mud.m3u8/WEB/1.0?ce=01c1a12d2a2a6542045919fa6cc161ebefed15c6e79a44b4eb43824c1122138da1e1c7e018d027a48d2a0b66fc12fae60555369c116c3cdd6675b030cd491a6c9baf2a4782abbbc87bcd5d1f5ea751003ab11a7b529ef64865f5388d7fe43b416245059b217b32a9c00f817ad7f2e6d9358092e6438bc892&pf=3&uk=49989f549719922ef5446b540d47498a&rx=10658&expire=1676686725838&ip=142.188.84.231&sign=837be7265bd8e15c7cb68326dd0fe282&_ts=1676661525838",
	}
)

const (
	resultsDir string = "results"
	mpegDir    string = resultsDir + "/mpeg"
	tsDir      string = resultsDir + "/ts"

	mpegExt string = ".m3u8"
	tsExt   string = ".ts"
)

var stopper bool = true

func Start() error {
	if stopper {
		fmt.Println("FATAL: cannot continue because .ts results need to be padded with leading zeroes!")
		fmt.Println("This is a bug that was left unfixed previously")
		return nil
	}

	ctx := context.Background()

	if err := createStorageDirectory(ctx); err != nil {
		return err
	}
	if err := fetchAndStoreMPEGFiles(ctx); err != nil {
		return err
	}
	if err := fetchAndStoreTSFiles(ctx); err != nil {
		return err
	}

	return nil
}

func createStorageDirectory(ctx context.Context) error {
	if _, err := os.Stat(resultsDir); os.IsNotExist(err) {
		if err := os.Mkdir(resultsDir, os.ModePerm); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	return nil
}

func fetchAndStoreMPEGFiles(ctx context.Context) error {
	ok, err := createMPEGDir(ctx)
	if err != nil {
		return err
	}
	if ok {
		return nil
	}

	for i, endpoint := range season1Episodes {
		if err := fetchAndStoreMPEGFile(ctx, i+1, endpoint); err != nil {
			continue
		}
	}

	return nil
}

func fetchAndStoreMPEGFile(ctx context.Context, i int, endpoint string) (err error) {
	log.Println("Fetching MPEG file", i, "for", endpoint)

	defer func() {
		if err != nil {
			log.Println("Errored fetching MPEG file", i, "for", endpoint)
		}
		log.Println("Done fetching MPEG file", i, "for", endpoint)
	}()

	var req *http.Request
	req, err = http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	var f *os.File
	f, err = os.CreateTemp(os.TempDir(), "*")
	if err != nil {
		return
	}

	if _, err = io.Copy(f, resp.Body); err != nil {
		return
	}
	resp.Body.Close()
	f.Close()

	if err = os.Rename(f.Name(), filepath.Join(mpegDir, strconv.Itoa(i)+mpegExt)); err != nil {
		return
	}

	return
}

func fetchAndStoreTSFiles(ctx context.Context) error {
	ok, err := createTSDir(ctx)
	if err != nil {
		return err
	}
	if ok {
		return nil
	}

	for i := range season1Episodes {
		endpoints, err := parseMPEGFile(ctx, i+1)
		if err != nil {
			return err
		}

		for j, endpoint := range endpoints {
			if err := fetchAndStoreTSFile(ctx, i+1, j+1, endpoint); err != nil {
				return err
			}
		}
	}

	return nil
}

func parseMPEGFile(ctx context.Context, i int) ([]string, error) {
	b, err := os.ReadFile(filepath.Join(mpegDir, strconv.Itoa(i)+mpegExt))
	if err != nil {
		return nil, nil
	}

	endpoints := make([]string, 0)
	lines := bytes.Split(b, []byte{'\n'})
	for _, line := range lines {
		if len(line) < 8 {
			continue
		}

		if line[0] != 'h' {
			continue
		}
		if line[1] != 't' {
			continue
		}
		if line[2] != 't' {
			continue
		}
		if line[3] != 'p' {
			continue
		}
		if line[4] != 's' {
			continue
		}
		if line[5] != ':' {
			continue
		}
		if line[6] != '/' {
			continue
		}
		if line[7] != '/' {
			continue
		}

		endpoints = append(endpoints, string(line))
	}

	return endpoints, nil
}

func fetchAndStoreTSFile(ctx context.Context, episode, file int, endpoint string) (err error) {
	log.Println("Fetching TS file for episode", episode, "file", file, "at", endpoint)

	defer func() {
		if err != nil {
			log.Println("Errored fetching TS file for episode", episode, "file", file, "at", endpoint)
		}
		log.Println("Done fetching TS file for episode", episode, "file", file, "at", endpoint)
	}()

	var req *http.Request
	req, err = http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	var f *os.File
	f, err = os.CreateTemp(os.TempDir(), "*")
	if err != nil {
		return
	}

	if _, err = io.Copy(f, resp.Body); err != nil {
		return
	}
	resp.Body.Close()
	f.Close()

	if err = os.Rename(f.Name(), filepath.Join(tsDir, strconv.Itoa(episode)+"-"+strconv.Itoa(file)+tsExt)); err != nil {
		return
	}

	return
}

func createMPEGDir(ctx context.Context) (bool, error) {
	if _, err := os.Stat(mpegDir); os.IsNotExist(err) {
		if err := os.Mkdir(mpegDir, os.ModePerm); err != nil {
			return false, err
		}
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

func createTSDir(ctx context.Context) (bool, error) {
	if _, err := os.Stat(tsDir); os.IsNotExist(err) {
		if err := os.Mkdir(tsDir, os.ModePerm); err != nil {
			return false, err
		}
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}
