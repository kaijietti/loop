# NOTE:
# Must not run this script directly, follow the timeline t0-t11 manually.

# t0 Bootstrap fluentd to collect container logs. [in terminal#1]
docker run -it -p 24224:24224 -v ./test.conf:/fluentd/etc/test.conf -e FLUENTD_CONF=test.conf fluent/fluentd:latest

# t1 Build and Run this program called 'loop' [in another terminal#2]
docker build -t loop:latest .
docker run -it --log-driver=fluentd --log-opt fluentd-address=127.0.0.1:24224 loop:latest

# t5 Now back to terminal#1 and stop the fluentd server

# t6 Observing terminal#2 for about half a minute, you will note that the terminal stop logging 'Hello' message.

# t10 Now try to stop this container or just press CTRL+C in terminal#2, you will note that container is stuck.
# t11 You will fail to stop or rm this container, unless restart docker.service.