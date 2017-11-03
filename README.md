
export YOUR_API_KEY=xxxxxxxxxx
curl -X POST "https://language.googleapis.com/v1/documents:analyzeSentiment?key=$YOUR_API_KEY" -H "Content-Type: application/json" --data-binary @postdata.json > result.json