using System;

class Program
{
  static void Main()
  {
    string s = Console.ReadLine();
    // s = "abracadabra"
    string[] input = Console.ReadLine().Trim().Split();
    int i = int.Parse(input[0]) - 1;
    int j = int.Parse(input[1]) - 1;

    Console.WriteLine(j - i + 2);
    string s1 = "";
    for (; i < j; i++)
    {
      s1 += s[i];
      Console.WriteLine(s1);
    }
  }
}
