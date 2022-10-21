interface Transaction {
    id: string,
    uid?: string,
    item_id: string,
    category?: string[],
    location?: string,
    name?: string,
    amount: number,
    iso_currency_code: string,
    date: Date,
    pending: boolean,
    merchant_name?: string,
    payment_channel?: string,
    shared_with?: string,
    split_amount?: string,
};

function ConvertTransaction(o: any) {
    const tx: Transaction = {
        id: o.transaction_id,
        item_id: o.item_id,
        pending: o.pending,
        merchant_name: o.merchant_name,
        amount: o.amount,
        date: o.date,
        category: o.category,
        iso_currency_code: o.iso_currency_code, 
    }
    return tx;
}

export { Transaction, ConvertTransaction };