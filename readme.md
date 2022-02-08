We will have a user structure consisting of fields like name,username,passsword,role/designation,BellaLevel,BibaLevel,RSAprivateKey,RSAPublicKey and slice of mails received.
According to the designation of the user, the BellaLevel and BibaLevel will be assigned.

We will also have a file structure with fields as name,BellaLevel and BibaLevel.

New users can be created and their passwords will stored after salting and hashing.
Menu consists of Login,Register,Logout.
While logging in,we will compare hashed+salted passwords.

After login - Read from file | Write into file | Send private messages to a user(Email) |See inbox mails
Read from a file
The user can read according to Bella and Biba levels.
Write into file
The user write into files according to Bella and Biba levels

Send private msg
Sender can see a list of users along with their public keys.
The user can send mails to a particular user by encrypting the msg with the receiver's public key and appending the MAC code with it.
At the reciever's end,the MAC will be authentication and decrypted using his/her own private key
While sending the mail, the mail body will consist of encrypted(Sender name + msg) 
Now mail received will be in text document at reciever's side.

See inbox mails
User can see his/her inbox i.e the mails he recevied.


Log book
Login/Logout details
File access -read/write/create details
Mail sent details (who sent mail to whom)
