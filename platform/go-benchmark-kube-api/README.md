# Go benchmark kube-api

As a little exercise to start learning go, I've created a little tool to benchmark my kube api server.
For now it's fairly simple and contains the following functionality:
- [x] Select the number of requests per second
- [x] Select total execution time
- [x] Calculate the number of failed requests
- [x] Calculate the number of requests that throttle based on an abnormally long request time (>1s)
- [x] Calculate the average query time

We use go routines spaced by a ticker to launch the correct number of requests. 
The tool uses the default kubeconfig and is launched with the following command:
```bash
$: go run benchmark.go -rps 1 -duration 120
Benchmark finished!
Total number of requests sent: 120
Total number of failed requests: 0
Number of requests with throttling detected: 0
Average response time: 0.015539
```
For the moment, the results are not relevant because I'm blocked by a rate limit on the client side.
```bash
$: go run benchmark.go -rps 4 -duration 5
I0922 18:04:09.874384   43487 request.go:700] Waited for 1.000128103s due to client-side throttling, not priority and fairness, request: GET:https://<LB_IP>:6443/api/v1/nodes
Benchmark finished!
Total number of requests sent: 20
Total number of failed requests: 0
Number of requests with throttling detected: 1
Average response time: 0.209232
```
A good configuration of the client will allow me to bypass it.