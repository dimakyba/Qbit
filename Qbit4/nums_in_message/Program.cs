using System;

class Program
{
  static void Main(string[] args)
  {
    int n = int.Parse(Console.ReadLine());
    string s = Console.ReadLine();
    int counter = 0;
    int sum = 0;

    foreach (char c in s)
    {
      if (char.IsDigit(c))
      {
        sum += (int)(c - '0');
        counter++;
      }
    }

    Console.WriteLine(counter);
    Console.WriteLine(sum);
  }
}
