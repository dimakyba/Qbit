using System;

namespace Program
{
    class Program
    {
        static void Main(string[] args)
        {
            string[] data = Console.ReadLine().Trim().Split();
            double x1 = double.Parse(data[0]);
            double y1 = double.Parse(data[1]);
            double x2 = double.Parse(data[2]);
            double y2 = double.Parse(data[3]);
            Console.WriteLine(Math.Sqrt(Math.Pow((x1-x2),2) + Math.Pow((y1-y2), 2)));

        }
    }
}
