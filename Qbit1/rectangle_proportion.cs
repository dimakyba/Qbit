using System;

namespace Program
{
    class Program
    {
        static void Main(string[] args)
        {
            string[] data = Console.ReadLine().Trim().Split();
            int p = Convert.ToInt32(data[0]);
            int x = Convert.ToInt32(data[1]);
            int y = Convert.ToInt32(data[2]);
            float temp = (float)p / (2.0f * (x + y));
            float a = x * temp;
            float b = y * temp;
            Console.WriteLine(Math.Round(a * b, 6));

        }
    }
}
