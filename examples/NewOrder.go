package main

// WARNING: THIS IS EXPERIMENTAL AND NON-WORKING DRAFT
// WARNING: THIS IS EXPERIMENTAL AND NON-WORKING DRAFT
// WARNING: THIS IS EXPERIMENTAL AND NON-WORKING DRAFT
// WARNING: THIS IS EXPERIMENTAL AND NON-WORKING DRAFT
// WARNING: THIS IS EXPERIMENTAL AND NON-WORKING DRAFT

import (
    "context"
    "fmt"

    "github.com/davecgh/go-spew/spew"
    serumgo "github.com/gagliardetto/serum-go"
    "github.com/gagliardetto/solana-go"
    "github.com/gagliardetto/solana-go/rpc"
    confirm "github.com/gagliardetto/solana-go/rpc/sendAndConfirmTransaction"
    "github.com/gagliardetto/solana-go/rpc/ws"
)

func main() {
    // WARNING: THIS IS EXPERIMENTAL AND NON-WORKING DRAFT
    // WARNING: THIS IS EXPERIMENTAL AND NON-WORKING DRAFT
    // WARNING: THIS IS EXPERIMENTAL AND NON-WORKING DRAFT
    // WARNING: THIS IS EXPERIMENTAL AND NON-WORKING DRAFT
    // WARNING: THIS IS EXPERIMENTAL AND NON-WORKING DRAFT
    // WARNING: THIS IS EXPERIMENTAL AND NON-WORKING DRAFT
    // WARNING: THIS IS EXPERIMENTAL AND NON-WORKING DRAFT
    serumgo.SetProgramID(serumgo.SerumDexV1)
    // Create a new RPC client (TODO: you need to select the appropriate network):
    cluster := rpc.DevNet
    rpcClient := rpc.New(cluster.RPC)

    // Create a new WS client (used for confirming transactions).
    // NOTE: here too you need to select the appropriate net.
    wsClient, err := ws.Connect(context.Background(), cluster.WS)
    if err != nil {
        panic(err)
    }

    // Load the account that will pay for the transaction and can sign stuff:
    owner, err := solana.PrivateKeyFromSolanaKeygenFile("/path/to/.config/solana/id.json")
    if err != nil {
        panic(err)
    }
    fmt.Println("owner public key:", owner.PublicKey())

    recent, err := rpcClient.GetRecentBlockhash(context.TODO(), rpc.CommitmentFinalized)
    if err != nil {
        panic(err)
    }
    instructions := []solana.Instruction{}
    signers := make([]solana.PrivateKey, 0)
    // NOTE: accounts.Owner is marked as signer, so it must be added to the signers.
    signers = append(signers, owner)
    {

        // Parameters:
        args := serumgo.NewOrderInstructionV1{
            Side:       serumgo.SideAsk,
            LimitPrice: 1720,
            MaxQty:     650000,
            OrderType:  serumgo.OrderTypeLimit,
            ClientId:   1608306862011613462,
        }
        // Accounts:
        market := solana.MustPublicKeyFromBase58("TODO")
        openOrders := solana.MustPublicKeyFromBase58("TODO")
        requestQueue := solana.MustPublicKeyFromBase58("TODO")
        orderPayer := solana.MustPublicKeyFromBase58("TODO")
        ownerAcc := solana.MustPublicKeyFromBase58("TODO")
        coinVault := solana.MustPublicKeyFromBase58("TODO")
        pcVault := solana.MustPublicKeyFromBase58("TODO")
        splTokenProgram := solana.MustPublicKeyFromBase58("TODO")
        rentSysvar := solana.MustPublicKeyFromBase58("TODO")
        feeDiscounts := solana.MustPublicKeyFromBase58("TODO")

        newOrderInstruction := serumgo.NewNewOrderInstructionBuilder().
            SetArgs(args).
            SetMarketAccount(market).
            SetOpenOrdersAccount(openOrders).
            SetRequestQueueAccount(requestQueue).
            SetOrderPayerAccount(orderPayer).
            SetOwnerAccount(ownerAcc).
            SetCoinVaultAccount(coinVault).
            SetPcVaultAccount(pcVault).
            SetSplTokenProgramAccount(splTokenProgram).
            SetRentSysvarAccount(rentSysvar).
            SetFeeDiscountsAccount(feeDiscounts)

        instructions = append(instructions, newOrderInstruction.Build())
    }

    tx, err := solana.NewTransaction(
        instructions,
        recent.Value.Blockhash,
        solana.TransactionPayer(owner.PublicKey()),
    )
    if err != nil {
        panic(err)
    }

    _, err = tx.Sign(
        func(key solana.PublicKey) *solana.PrivateKey {
            for _, candidate := range signers {
                if candidate.PublicKey().Equals(key) {
                    return &candidate
                }
            }
            return nil
        },
    )
    if err != nil {
        panic(fmt.Errorf("unable to sign transaction: %w", err))
    }
    // Pretty print the transaction:
    spew.Dump(tx)

    // Send transaction, and wait for confirmation:
    sig, err := confirm.SendAndConfirmTransaction(
        context.TODO(),
        rpcClient,
        wsClient,
        tx,
    )
    if err != nil {
        panic(err)
    }
    spew.Dump(sig)
}
