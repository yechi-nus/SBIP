package com.example.demo;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import org.hyperledger.fabric.gateway.Contract;
import org.hyperledger.fabric.gateway.ContractException;
import org.hyperledger.fabric.gateway.Gateway;
import org.hyperledger.fabric.gateway.Network;
import org.hyperledger.fabric.gateway.Wallet;
import org.hyperledger.fabric.gateway.Wallets;

import java.io.IOException;
import java.nio.ByteBuffer;
import java.nio.charset.StandardCharsets;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.concurrent.TimeoutException;

@SpringBootApplication
@RestController
public class DemoApplication {

	public static void main(String[] args) {
		SpringApplication.run(DemoApplication.class, args);
	}

	@GetMapping("/hello")
	public String hello(@RequestParam(value = "name", defaultValue = "World") String name) {
		return String.format("Hello %s!", name);
	}

	@GetMapping("/send")
	public SendResult send(@RequestParam(value = "id") int id, @RequestParam(value = "timestamp") String timestamp, @RequestParam(value = "from") String from, @RequestParam(value = "to") String to, @RequestParam(value = "amount") double amount, @RequestParam(value = "key") String key, @RequestParam(value = "signature") String signature) {
		if (id < 0 || id > 2000000000) {
			return new SendResult("Error id", id);
		}
		if (amount <= 0) {
			return new SendResult("Error amount: "+amount, id);
		}
		try {
			Path walletDirectory = Paths.get("wallet");
		    Wallet wallet = Wallets.newFileSystemWallet(walletDirectory);
		
		    // Path to a common connection profile describing the network.
		    Path networkConfigFile = Paths.get("connection.json");
		
		    // Configure the gateway connection used to access the network.
		    Gateway.Builder builder = Gateway.createBuilder()
		            .identity(wallet, "appUser")
		            .networkConfig(networkConfigFile);
		
		    // Create a gateway connection
		    try (Gateway gateway = builder.connect()) {
		
		        // Obtain a smart contract deployed on the network.
		        Network network = gateway.getNetwork("mychannel");
		        Contract contract = network.getContract("fabcoin");
		
		        // Submit transactions that store state to the ledger.
		        //byte[] createCarResult = contract.createTransaction("createCar")
		        //        .submit("CAR10", "VW", "Polo", "Grey", "Mary");
		        //System.out.println(new String(createCarResult, StandardCharsets.UTF_8));
		
		        // Evaluate transactions that query state from the ledger.
		        //byte[] queryAllCarsResult = contract.evaluateTransaction("queryAllCars");
		        //System.out.println(new String(queryAllCarsResult, StandardCharsets.UTF_8));
		        
		        byte[] querySend = contract.createTransaction("querySend").submit(""+id, timestamp, from, to, ""+amount, key, signature);
		        System.out.println(new String(querySend, StandardCharsets.UTF_8));
		        return new SendResult("OK", id);
		    } catch (Exception ex)
			{
				ex.printStackTrace();
				return new SendResult("Error:"+ex.toString(), id);
			}
	    } catch (Exception ex)
		{
			ex.printStackTrace();
			return new SendResult("Error:"+ex.toString(), id);
		}
	}
	
	@GetMapping("/get-balance")
	public GetBalanceResult getbalance(@RequestParam(value = "id") int id, @RequestParam(value = "timestamp") String timestamp, @RequestParam(value = "account") String account, @RequestParam(value = "key") String key, @RequestParam(value = "signature") String signature) {
		double amount = -1;
		if (id < 0 || id > 2000000000) {
			return new GetBalanceResult("Error id:", id, -1);
		}
		try {
			Path walletDirectory = Paths.get("wallet");
		    Wallet wallet = Wallets.newFileSystemWallet(walletDirectory);
		
		    // Path to a common connection profile describing the network.
		    Path networkConfigFile = Paths.get("connection.json");
		
		    // Configure the gateway connection used to access the network.
		    Gateway.Builder builder = Gateway.createBuilder()
		            .identity(wallet, "appUser")
		            .networkConfig(networkConfigFile);
		
		    // Create a gateway connection
		    try (Gateway gateway = builder.connect()) {
		
		        // Obtain a smart contract deployed on the network.
		        Network network = gateway.getNetwork("mychannel");
		        Contract contract = network.getContract("fabcoin");
		
		        byte[] queryGetBalance = contract.evaluateTransaction("getBalance", ""+id, timestamp, account, key, signature);
		        String balanceString = new String(queryGetBalance, StandardCharsets.UTF_8);
		        System.out.println("Balance:"+balanceString);
		        
		        amount = Double.parseDouble(balanceString);
		        if (amount >= 0) {
		        	return new GetBalanceResult("OK", id, amount);
		        }
		    } catch (Exception ex)
			{
				ex.printStackTrace();
				return new GetBalanceResult("Error:"+ex.toString(), id, amount);
			}
	    } catch (Exception ex)
		{
			ex.printStackTrace();
			return new GetBalanceResult("Error:"+ex.toString(), id, amount);
		}
		return new GetBalanceResult("Error", id, 0);
	}
}