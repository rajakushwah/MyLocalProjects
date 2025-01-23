class SimpleDB:

    def __init__(self):
        self.db = {}
        self.transactions =[]

    def set(self, key, value):
        """set sets the value associated with the key"""
        if self.transactions:
            if key not in self.transactions[-1]:
                self.transactions[-1][key] = self.db.get(key)
        self.db[key] = value

        # pass

    def get(self, key):
        """
        get returns the value associated with the key
        get should raise a KeyError if the key doesn't exist
        """
        if key not in self.db:
            raise KeyError(f"Key '{key}' does not exists")
        return self.db[key]


    def unset(self, key):
        """unset should delete the key from the db"""
        if key in self.db:
            if self.transactions:
                if key not in self.transactions[-1]:
                    self.transactions[-1][key] = self.db[key]
            del self.db[key]
        # pass

    def begin(self):
        """begin starts a new transaction"""
        self.transactions.append({})
        # pass

    def commit(self):
        """
        commit applies all transactions
        it should raise an Exception if there is no ongoing transaction
        """
        if not self.transactions:
            raise KeyError(f"no ongoing transaction")

        self.transactions.clear()
        # pass

    def rollback(self):
        """
        rollback undoes the most recent transaction
        it should raise an Exception if there is no ongoing transaction
        """
        if not self.transactions:
            raise KeyError(f"no ongoing transaction")
        current_transaction  = self.transactions.pop()
        for key , value in current_transaction.items():
            # print("rollback : key , value", key , value)
            if value is None:
                if key in self.db:
                    del self.db[key]
            else:
                self.db[key] = value

        # pass
