import json

data = json.loads(input())
if int(data["CurrentScore"]) > 500:
    print("b", end="")
else:
    print("r", end="")
