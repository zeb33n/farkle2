#include <cjson/cJSON.h>
#include <stdio.h>
#include <string.h>
#include <unistd.h>

#define BUFFERLEN 1024

int main() {
  char json[BUFFERLEN];
  json[BUFFERLEN - 1] = '\0';
  cJSON* root;
  int score;
  while (1) {
    read(STDIN_FILENO, &json, BUFFERLEN - 1);
    root = cJSON_Parse(json);
    score = cJSON_GetObjectItem(root, "CurrentScore")->valueint;
    if (score < 500) {
      printf("r");
      fflush(stdout);
    } else {
      printf("b");
      fflush(stdout);
    }
    memset(json, '\0', BUFFERLEN);
  }
  return 0;
}
