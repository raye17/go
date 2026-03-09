FROM swr.cn-north-4.myhuaweicloud.com/ddn-k8s/quay.io/devtron/nats:2.9.3-alpine

CMD sh -c 'i=0; while true; do echo "[TASK] running... count=$i"; i=$((i+1)); sleep 2; done'



docker run --rm \
  --name task-log-test \
  busybox sh -c '
  echo "=== TASK LOG TEST START ==="
  i=0
  while [ $i -lt 20 ]; do
    echo "TASK_LOG task_id=123 node_id=abc count=$i"
    i=$((i+1))
    sleep 1
  done
  echo "=== TASK LOG TEST END ==="
'
