package query

var (
	CreateReceipt = `
	 INSERT INTO
	   receipts(
		resep,
		bahan,
		kategori
	  )VALUES(
		:resep,
		:bahan,
		:kategori
	);
	`

	FindOneReceipt = `
    SELECT
    %v
    FROM receipts
    WHERE %v LIMIT 1
   `

	UpdateReceipt = `
   	UPDATE receipts
	SET %v
   	WHERE %v
   `

	FindAllReceipt = `
	SELECT * FROM receipts;
   `

	DeleteReceipt = `
	DELETE FROM receipts
	WHERE resep = ?;
   `
)
