## Project Summary:

Stand alone service calls CoinMarketcap API and updates the database with up to date data for each coin/token. 

- Built in go 1.23.5
- Uses John Barton's [dotenv](https://github.com/joho/godotenv)

### Setup/Run:

- Clone the repository
- Get an API key from [Coinmarketcap](https://coinmarketcap.com/api/)
- Rename .env.example to .env
- Add your API Key to .env CMC_API_KEY=
- In main.go, set the time interval for the API call based on how many credits you want to spend. Default set to 10 seconds (updater := time.NewTicker(10 * time.Second))
- Run in command prompt: 
    ```shell
    docker-compose up --build
    ```

