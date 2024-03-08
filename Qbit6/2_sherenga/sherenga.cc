#include <iostream>
#include <vector>
int main()
{
  int n, value, count = 1;
  std::cin >> n;
  std::vector<int> arr(n);
  for (int i = 0; i < n; i++)
  {
    std::cin >> arr[i];
  }
  std::cin >> value;
  for (int i = 0; i < n; i++)
  {
    if (arr[i] >= value)
    {
      count++;
    }
  }
  std::cout << count << "\n";
  return 0;
}
