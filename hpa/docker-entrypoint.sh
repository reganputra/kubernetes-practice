!#/bin/sh

if [ -z $NUM_THREADS]
then
 echo "[!] Running application with default  thread"
 openssl speed -multi 50 | tail /dev/zero
else
echo "[*] Running application with $NUM_THREADS threads"
    openssl speed -multi $NUM_THREADS | tail /dev/zero
fi