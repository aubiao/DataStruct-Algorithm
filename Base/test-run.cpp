//
// Created by aubiao ou on 2022/3/23.
//
#include <cassert>
//#include <ctime>
#include <dirent.h>
#include <iostream>
#include <unistd.h>
#include <unordered_map>
#include <unordered_set>

#include "leveldb/db.h"

#define KV_NUM_DEFAULT (50000)
#define KEY_SIZE (24)

using namespace std;
using namespace leveldb;

#define GEN_VALUE(v, i, size) \
  do {                        \
    memset(v, 2 + i, size);   \
  } while (0);

static int runCount = 0;
int KVNum = KV_NUM_DEFAULT;
int VALUE_SIZE = 10;
std::unordered_map<int, std::string> loadKeys;

void timer(bool start = true, const char *label = 0) {
    static std::chrono::system_clock::time_point startTime;
    if (start) {
        startTime = std::chrono::system_clock::now();
    } else {
        std::chrono::system_clock::time_point endTime =
                std::chrono::system_clock::now();
        printf("Elapsed Time (%s): %.6lf s\n", label,
               std::chrono::duration_cast<std::chrono::microseconds>(endTime -
                                                                     startTime)
                       .count() /
               1000.0 / 1000.0);
    }
}

void gen_key(char *k, int i, int j) {
    char c;
    do {
        c = rand() % (sizeof(char) << 8);
    } while (c == 0);
    memset(k + j, c, 1);
}

void testSetKey(leveldb::DB *db) {
    srand(123987);
    runCount += 1;

    int count = 0, failed = 0;
    char key[KEY_SIZE];
    std::unordered_set<int> usedKeys;
    string value;
    for (int w = 0; w < VALUE_SIZE; ++w) {
        value.push_back('a');
    }
    for (int i = 0; i < KVNum; i++) {
        int ki = -1;
        // always consume the key (in the random sequence)
        if (runCount == 1) {
            ki = i;
            for (int j = 0; j < KEY_SIZE; j++) {
                gen_key(key, i, j);
            }
            loadKeys[i] = string(key);
        } else {
            ki = i;
            memcpy(key, loadKeys.at(ki).c_str(), KEY_SIZE);
        }
        //    GEN_VALUE(value, ki + runCount, VALUE_SIZE);
        leveldb::Status status = db->Put(WriteOptions(), key, value);
        if (!status.ok()) {
            failed++;
        }
        count++;
    }

    printf(">>> Issued %d SET requests (%d failed)\n", count, failed);
}

void testReadBackKey(leveldb::DB *db) {
    srand(123987);

    char key[KEY_SIZE];

    int failed = 0;
    string value;
    for (int w = 0; w < VALUE_SIZE; ++w) {
        value.push_back('a');
    }
    for (int i = 0; i < KVNum; i++) {
        for (int j = 0; j < KEY_SIZE; j++) {
            gen_key(key, i, j);
        }
        string res;
        leveldb::Status status = db->Get(ReadOptions(), key, &res);
        if (res != value) failed++;
    }
    printf(">>> Issued %d GET requests (%d failed)\n", KVNum, failed);
    if (failed > 0) assert(0);
}

void removeDir(const char *dirPath) {
    struct dirent *dirP;
    DIR *dir = opendir(dirPath);
    while ((dirP = readdir(dir)) != nullptr) {
        string name = dirP->d_name;
        string path = dirPath;
        if (dirP->d_type == DT_REG) {
            // 文件
            remove((path + name).c_str());
        } else if (dirP->d_type == DT_DIR) {
            // 文件夹
            // 此处需注意特殊系统文件
            if (name != "." && name != "..") {
                removeDir((path + name + "/").c_str());
            }
        }
    }
    remove(dirPath);
}

int main(int argc, char **argv) {
    if (argc >= 3) {
        KVNum = atoi(argv[1]);
        VALUE_SIZE = atoi(argv[2]);
    }

    leveldb::DB *db;
    leveldb::Options options;
    options.create_if_missing = true;
    leveldb::Status status = leveldb::DB::Open(options, "testdb", &db);
    assert(status.ok());

    fprintf(stderr, "> Beginning of tests (KVNum=%d,VALUE_SIZE=%d)\n", KVNum,
            VALUE_SIZE);
    timer(true);
    testSetKey(db);
    timer(false, "SET");
    string property;
    db->GetProperty("leveldb.sstables", &property);
    db->GetProperty("leveldb.stats", &property);
    timer(true);
    testReadBackKey(db);
    timer(false, "GET");
    const char *dirPath = "./testdb/";
    removeDir(dirPath);

    delete db;
    return 0;
}
