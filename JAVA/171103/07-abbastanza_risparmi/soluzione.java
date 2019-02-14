import java.util.*;
public class soluzione {
	public static void main(String[]args){
		Scanner in = new Scanner(System.in);
		int day = 1;
		int sum = 0;
		int price = in.nextInt();
		int i = in.nextInt();
		
		while (i != 0) {
			sum = sum + i;
			if (sum < price)
				day++;
			i = in.nextInt();
		}
			
		if (sum < price)
			day = 0;
			
		System.out.println(day);
	}
}
