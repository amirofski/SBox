# SBox
How To run Dev mode? <br/>
```ignite chain serve```
Build<br />
```ignite scaffold chain build```
Define Messages <br />
```ignite scaffold message storeData data:string owner:string```<br />
```ignite scaffold message retrieveData id:uint```
### Send Transactions
## Send a transaction to store data:<be />
``` [SBox]d tx storage storeData "data to store" --from alice```
