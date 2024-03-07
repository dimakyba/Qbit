using System;

namespace Program
{
    class Program
    {
        static void Main(string[] args)
        {
            string[] data = Console.ReadLine().Trim().Split();
            uint a = Convert.ToUInt16(data[0]);
            uint b = Convert.ToUInt16(data[1]);
            uint c = Convert.ToUInt16(data[2]);
            string[] data2 = Console.ReadLine().Trim().Split();
            uint h = Convert.ToUInt16(data2[0]);
            uint l = Convert.ToUInt16(data2[1]);
            double area = (a + b) * 2 * c + (a*b);
            double rulonArea = h * l / 1e6;
            Console.WriteLine((int)Math.Ceiling(area * 0.85 / rulonArea * 1.1));
        }
    }
}
