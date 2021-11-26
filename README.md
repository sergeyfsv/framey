# Intro
go.framey.io/speedtest library is created for measuring an internet connection speed by using multiple services: speedtest.net and fast.com

# Usage

* Import library "go get -u go.framey.io/speedtest
* Do imports
```
	. "go.framey.io/speedtest"
	"go.framey.io/speedtest/internal/types"
```
* Create a new measurement configuration
```
measurementConfiguration := types.DefaultConfiguration()
```
* Adjust default timeout settings if necessary
  * **ConfigTimeout**:    Timeout for getting fast.com/speed.net configuration before a test. Default: 5 seconds,
  * **DownloadTimeout**:  Timeout for an each download operation in a thread pool. Default: 10 seconds,
  * **UploadTimeout**:    Timeout for an each upload operation in a thread pool. Default: 10 seconds,

* Adjust the result output format
    * **MeasurementRatio**: **1**. Setting it to 1 will return speed in Bytes per Second
    * **MeasurementRatio**: **8 / float64(types.Mbps)**. Converts it to Mbits per Second
* Create a new measurement by providing measurement endpoint and the configuration
  * To test speed with speedtest.net
    ```
    measurement := NewMeasurement(types.SpeedtestdotnetEndpoint, &measurementConfiguration)
    ```
  * To test speed with fast.com
    ```
    measurement := NewMeasurement(types.FastdotcomEndpoint, &measurementConfiguration)
    ```
* Call either the Upload or Download test
```
	speed, err := Upload(&measurement)
	speed, err := Download(&measurement)
```
  * **speed** is normalized according to the specified configuration.MeasurementRatio

# Makefile
The makefilie provided with the library allows to automate some routine tasks:
* Running tests
  ```
  make test
  ```
* Measuring tests coverage
  ```
  make test-coverage
  ```

# Things to improve
  * It is possible to provide measurementCallback core.MeasurementCallback in the public library's API also
  * For the sake of the interface unification all servers's speed measurement initialization have been wrapped with Upload/Download method. It has own performance issue since for each Upload/Download execution it either get Fast.com token or servers for speedtest.net
  * It might make sense to extract a parameter that controls the number of attempts to measure the speed. This parameter is already there in the library.
  * Create more unit tests. However, since the library is based on another one internally, that is not the obvious task because of its internal architecture. So, such tests will mainly be mock-based tests.
# Benchmarks
  I do not see any reasons with the current structure of library providing any benchmarks, since that its code is only about initializing and calling other methods.