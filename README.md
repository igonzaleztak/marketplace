<h1>Blockchain Marketplace</h1>
This is a Blockchain based platform that enables the purchase of IoT measurements directly from the IoT suppliers in an easy and transparent way. IoT provider can use this platform to publish their measurements, so clients can purchase them. The inherent properties of Blockchain technology allows the recording of every action produced within the platform. Thus, ensuring transparency.

<h2>Architecure</h2>
In this <a href="#gen-arch">figure</a> you can see the general architecture of the system. The core of the platform is the Blockchain. This element is an Ethereum Blockchan that uses Clique as consensus model. Inside the Blockchain there are four smart contracts running on it.

<ul><a href="./storage/contracts/dataContract/data.sol">Data SC</a>: Stores the encrypted URL where the measurement is located in the Blockchain, so it is available to be purchased. In addition to the URL,  this contract also stores the symmetric key, encrypted with the public key of the administrator of the platform,  to decipher the value of the measurement </ul>
<ul><a href="./storage/contracts/accessContract/accessContract.sol">Access Control SC</a>: Controls the users who can introduce measurements in the Blockchain. Only those IoT suppliers approved by the administrator of the platform can insert new content in the platform. Moreover, this SC stores the public keys of the customers that participates in the platform.</ul>
<ul><a href="./storage/contracts/balanceContract/balance.sol">Payment SC</a>: Manages the purchases of the clients. It guarantees that a measurement has been bought by a client</ul>
<ul><a href="./storage/contracts/balanceContract/ERC20.sol">ERC20 SC</a>: Defines the Token that will be used in the platform and enables its exchange between customers and IoT suppliers.</ul>

These are ruled by the administrator of the platform. This user is responsible for adding or/and removing new IoT producers to the platform, delivering the measurements purchased by customers and configuring the parameters of the smart contract (i.e. total amount of tokens).

Besides the Blockchain, several elements participate within the platform. One of these components is the IoT connector. This is a proxy between the IoT gateway and the platform. It encrypts the measurements with a random symmetric key and stores them in the IPFS network, then it stores all the information needed to retrieve the measurements in the Blockchain. Specifically, it stores the IPFS URL and the symmetric key encrypted with the public key of the administrator of the platform., it stores the IPFS URL and the symmetric key encrypted with the public key of the administrator of the platform.

To store the value of the measurements, the platform uses a private IPFS network. Each of the IoT suppliers and the marketplace has an IPFS node that they can use to store the value of the measurements. This node is protected with an authentication module that controls who inserts information in the network. To do so, it uses the Blockchain, in particular the Access SC, to check whether the IoT provider has access to the IPFS network or not. Besides the IoT IPFS nodes, the marketplace has its own IPFS node to retrieve the value of the measurements.

The marketplace is in charge of listening to the purchase events produced in the Blockchain whenever a client purchases a measurement. Once this component listens to an event, it retrieves the complete value of the acquired measurement from the IPFS storage and deciphers the symmetric key stored in the Data SC with its private key. Then, it encrypts the symmetric key with the public key of the client who made the purchase and sends the encrypted measurement plus the encrypted symmetric key to the client. The customer only has to decipher the symmetric key with its private key and use it to decrypt the value of the measurement. 

Customers can use a webserver (Clients Wallet) to browse and purchase the different measurements available in the platform. Through this component, clients can purchase measurements and see their value easily. 

<p align="center">
  <img src="docs/images/gen-arch.png" height="400px" width="700px" alt="Image">
  <p align="center" id="gen-arch">General architecture of the platform</p>
</p>
