import java.util.*;
public class soluzione {
	public static void main(String[]args) {
		Scanner in = new Scanner(System.in);
		
		int min = in.nextInt();
		//int min = Integer.MAX_VALUE;
		int x = 1;
		
		while (x > 0){
			x = in.nextInt();
		
		if (min > x & x != 0)
			min = x;
		
		if (x==0)
			System.out.print(min);
		}
		
	}
}
