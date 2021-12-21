package com.example.demo;

public class GetBalanceResult {
	public String code;
	public int id;
	public double amount;
	
	public GetBalanceResult(String code, int id, double amount) {
		this.id = id;
		this.code = code;
		this.amount = amount;
	}
}
