using System;

class Program
{
  static void Main()
  {
    string inputString = Console.ReadLine();
    int k = int.Parse(Console.ReadLine());
    int counter = 0;

    for (int i = 0; i < inputString.Length; i++)
    {
      if (inputString[i] == ' ')
      {
        counter++;

        if (counter == k)
        {
          Console.WriteLine(i + 1);
          return;
        }
      }
    }
    Console.WriteLine(0);
  }
}
