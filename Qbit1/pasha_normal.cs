using System;

class Program
{
  static void Main(string[] args)
  {
    // const int sumOfFive = 15;
    string[] data = Console.ReadLine().Trim().Split();
    int a = Int32.Parse(data[0];)
    int b = Int32.Parse(data[1];)
    int c = Int32.Parse(data[2];)
    int d = Int32.Parse(data[3];)
    // int result = sumOfFive - a - b - c - d;
    int result = 15 - a - b - c - d;
    Console.WriteLine(result);
  }
}
