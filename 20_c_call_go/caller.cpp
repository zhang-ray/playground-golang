// 使用方法：
// g++ -std=c++11 -o caller caller.cpp yigeku.so && ./caller

extern "C"{
#include <stdio.h>
#include "yigeku.h"
}
#include <iostream>

constexpr size_t theSize = 6;

int main() {
    GoInt a = 12;
    GoInt b = 99;
    printf("awesome.Sub(12,99) = %d\n", Sub(a, b));
    printf("awesome.Sine(1) = %f\n", (float)(Sine(1.0)));
    
    GoInt data[theSize] = {77, 12, 5, 99, 28, 23};
    GoSlice nums = {data, theSize, theSize};
    Sort(nums);
    
    for (int i = 0; i < theSize; i++){
        printf("%d\t", ((GoInt *)nums.data)[i]);
    }
    std::cout << std::endl;

    GoString msg = {"Hello from C!", 13};
    Log(msg);
}
