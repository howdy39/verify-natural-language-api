
# 準備
```
dep ensure
```

# 実行
```
go run template/main.go
```

# 結果
```
analyzeEntities
{"entities":[{"name":"Natural Language API","type":7,"metadata":{"mid":"/m/059nn","wikipedia_url":"https://en.wikipedia.org/wiki/Natural_language"},"salience":1,"mentions":[{"text":{"content":"Natural Language API"},"type":1}]}],"language":"ja"}
analyzeSentiment
{"document_sentiment":{"magnitude":0.8,"score":0.8},"language":"ja","sentences":[{"text":{"content":"Natural Language API を使うのは楽しいです"},"sentiment":{"magnitude":0.8,"score":0.8}}]}
analyzeSyntax
{"sentences":[{"text":{"content":"Natural Language API を使うのは楽しいです"}}],"tokens":[{"text":{"content":"Natural"},"part_of_speech":{"tag":12,"proper":1},"dependency_edge":{"head_token_index":2,"label":26},"lemma":"Natural"},{"text":{"content":"Language","begin_offset":8},"part_of_speech":{"tag":6,"proper":1},"dependency_edge":{"head_token_index":2,"label":26},"lemma":"Language"},{"text":{"content":"API","begin_offset":17},"part_of_speech":{"tag":6,"proper":1},"dependency_edge":{"head_token_index":4,"label":18},"lemma":"API"},{"text":{"content":"を","begin_offset":21},"part_of_speech":{"tag":9,"case":1,"proper":2},"dependency_edge":{"head_token_index":2,"label":45},"lemma":"を"},{"text":{"content":"使う","begin_offset":24},"part_of_speech":{"tag":11,"form":1,"proper":2},"dependency_edge":{"head_token_index":5,"label":71},"lemma":"使う"},{"text":{"content":"の","begin_offset":30},"part_of_speech":{"tag":9,"proper":2},"dependency_edge":{"head_token_index":7,"label":72},"lemma":"の"},{"text":{"content":"は","begin_offset":33},"part_of_speech":{"tag":9,"proper":2},"dependency_edge":{"head_token_index":5,"label":45},"lemma":"は"},{"text":{"content":"楽しい","begin_offset":36},"part_of_speech":{"tag":1,"form":4,"proper":2},"dependency_edge":{"head_token_index":7,"label":54},"lemma":"楽しい"},{"text":{"content":"です","begin_offset":45},"part_of_speech":{"ta
```