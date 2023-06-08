# Loop

> Forgive my naming `loop` which has nothing do with this bug.

Follow the steps by timeline (from t0 to t11) and 
this can reproduce a strange phenomenon that 
docker cannot stop/rm/logs/... a container (which run with fluentd log diver`async disable` and then failed to connect to a fluentd server)
**unless you restart the `docker.service`**.