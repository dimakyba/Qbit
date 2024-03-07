using System;

class Program
{
  static void Main()
  {
    string inputString = Console.ReadLine();
    string[] data = Console.ReadLine().Trim().Split();
    int i = int.Parse(data[0]);
    int j = int.Parse(data[1]);
    int substringLength = j - i + 1;
    string substring = inputString.Substring(i - 1, substringLength);
    Console.WriteLine(substringLength);
    Console.WriteLine(substring);
  }
}
