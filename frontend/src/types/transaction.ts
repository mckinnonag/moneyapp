interface Transaction {
    id: string,
    merchant_name: string,
    amount: number,
    date: Date,
    category: string[],
    iso_currency_code: string,
};

function ConvertTransaction(o: any) {
    const tx: Transaction = {
        id: o.transaction_id,
        merchant_name: o.merchant_name,
        amount: o.amount,
        date: o.date,
        category: o.category,
        iso_currency_code: o.iso_currency_code, 
    }
    return tx;
}

export { Transaction, ConvertTransaction };