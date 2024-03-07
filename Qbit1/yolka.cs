using System;

class Program
{
  static void Main()
  {
    string[] data = Console.ReadLine().Trim().Split();
    int shar = Int32.Parse(data[0]);
    int sos = Int32.Parse(data[1]);
    int sharBoxes = 0;
    int sosBoxes = 0;

    sharBoxes = shar / 3;
    shar = shar % 3;
    sosBoxes = sos / 4;
    sos = sos % 4;

    if (shar > 0)
      sharBoxes++;
    if (sos > 0)
      sosBoxes++;

    Console.WriteLine(sharBoxes + sosBoxes);
  }
}
