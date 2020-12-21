#include <iostream>
#include <map>


using namespace std;

int binarySearch(const int l[], int size, int target);
int fibonacci(int n, map<int, int> values);

int main() {

    int foo [] = {  1, 3, 4, 4, 6, 17, 79, 81, 90};



    cout << binarySearch(foo, size(foo), 79) <<endl;

    map<int, int> vals = map<int, int>();
    cout << fibonacci(100, vals) << endl;

    return 0;
}

int fibonacci(int n, map<int, int> values){
    if (n == 0 || n == 1){
        return n;
    }

    auto iter = values.find(n);
    if(iter == values.end()){
        return values[n] = fibonacci(n-1, values)+fibonacci(n-2, values);
    }else{
        return iter->second;
    }

}

int binarySearch(const int l[], int size, int target){
    int left = 0, right = size-1, mid;
    while (left < right){
        mid = (left+right)/2;
        if (l[mid] < target){
            left = mid+1;
        }else if (l[mid] > target){
            right = mid-1;
        }else{
            return mid;
        }
    }

    return -1;
}
