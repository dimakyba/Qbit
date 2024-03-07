using System;

class Program
{
  static void Main()
  {
    int n = int.Parse(Console.ReadLine());
    string message = Console.ReadLine();

    char mostFrequentChar = FindMostFrequentChar(message);

    Console.WriteLine((int)mostFrequentChar);
    Console.ReadLine();
  }

  static char FindMostFrequentChar(string message)
  {
    int[] charFrequency = new int[256];

    foreach (char c in message)
    {
      int index = (int)c;
      charFrequency[index]++;
    }

    char mostFrequentChar = ' ';
    int maxFrequency = 0;

    for (int i = 0; i < 256; i++)
    {
      if (charFrequency[i] > maxFrequency || (charFrequency[i] == maxFrequency && (char)i < mostFrequentChar))
      {
        mostFrequentChar = (char)i;
        maxFrequency = charFrequency[i];
      }
    }

    return mostFrequentChar;
  }
}
