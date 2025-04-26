#include <bitset>/stdc++.H>

using namespace std;

int main()
{

    int n = 10;
    int string = "";
    while (n > 0)
    {
        int bit = n % 2;
        string = string + char(bit); //-->"0101"
        n = n / 2; //--> 5 --> 2 -->1
    }

    reverse(str);
}