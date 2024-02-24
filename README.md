## Start
./start.sh

## Clean
./clean.sh

### create withdraw
curl --location --request POST 'http://127.0.0.1:8000/withdraw' \
--header 'Content-Type: application/json' \
--data '{
    "amount":400000,
    "recipient":"0x25e8F13860B125ea79B57C3AfF3782D1dAa4caa9"
}'


### confirm withdraw by manager
curl --location --request POST 'http://127.0.0.1:8000/confirm/1' \
--header 'Content-Type: application/json' \
--data '{
    "manager_id":1
}'

### get withdraw
curl --location --request GET 'http://127.0.0.1:8000/withdraw/1' 


