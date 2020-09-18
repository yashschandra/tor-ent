# start tor
tor &

while true
do
    # check if tor has created our hidden service directory
    if [ -d "/var/lib/tor/hidden_service" ]; then
        break
    fi
    # try again
    echo "finding directory in 5s"
    sleep 5
done

# print onion address
echo $(cat /var/lib/tor/hidden_service/hostname)

# starting golang server at 17666 to handle incoming requests
echo "starting golang server..."
cd /opt/go && go run *.go