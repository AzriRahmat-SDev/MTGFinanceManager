import React, { useState, Fragment } from 'react';
import { nanoid } from 'nanoid';
import data from './mock-data.json';
import ReadOnlyRow from './ReadOnlyRow';
import EditableRow from './EditableRow';

function Table() {
	const [cardDetails, setCardDetails] = useState(data);
	const [addFormData, setAddFormData] = useState({
		quantity: '',
		cardName: '',
		version: '',
		condition: '',
		language: '',
		finish: '',
		cardPrice: '',
	});

	const [editFormData, setEditFormData] = useState({
		quantity: '',
		cardName: '',
		version: '',
		condition: '',
		language: '',
		finish: '',
		cardPrice: '',
	});

	const [editCardId, setEditCardId] = useState(null);

	//important piece of code that should be looked at when handling events
	//records down every keystroke of the user in the "add card" component
	const handleAddFormChange = (e) => {
		e.preventDefault();

		const fieldName = e.target.getAttribute('name');
		const fieldValue = e.target.value;

		const newFormData = { ...addFormData };
		newFormData[fieldName] = fieldValue;

		setAddFormData(newFormData);
	};

	const handleEditFormChange = (e) => {
		e.preventDefault();

		const fieldName = e.target.getAttribute('name');
		const fieldValue = e.target.value;

		const newFormData = { ...editFormData };
		newFormData[fieldName] = fieldValue;

		setEditFormData(newFormData);
	};

	//Add card button of the "Add new card portion of the app"
	const handleAddFormSubmit = (e) => {
		e.preventDefault();

		const newCardDetail = {
			id: nanoid(),
			quantity: addFormData.quantity,
			cardName: addFormData.cardName,
			version: addFormData.version,
			condition: addFormData.condition,
			language: addFormData.language,
			finish: addFormData.finish,
			cardPrice: '$' + addFormData.cardPrice,
		};

		const newCardDetails = [...cardDetails, newCardDetail];
		setCardDetails(newCardDetails);
	};

	//submit inline edit for the rows upon clicking the save button
	const handleEditFormSubmit = (e) => {
		e.preventDefault();

		const editedCardDetail = {
			id: editCardId,
			quantity: editFormData.quantity,
			cardName: editFormData.cardName,
			version: editFormData.version,
			condition: editFormData.condition,
			language: editFormData.language,
			finish: editFormData.finish,
			cardPrice: '$' + editFormData.cardPrice,
		};

		const newCardDetails = [...cardDetails];

		const index = cardDetails.findIndex(
			(cardDetail) => cardDetail.id === editCardId
		);

		newCardDetails[index] = editedCardDetail;
		setCardDetails(newCardDetails);
		setEditCardId(null);
	};

	const handleEditClick = (e, cardDetails) => {
		e.preventDefault();
		setEditCardId(cardDetails.id);

		const formValues = {
			quantity: cardDetails.quantity,
			cardName: cardDetails.cardName,
			version: cardDetails.version,
			condition: cardDetails.condition,
			language: cardDetails.language,
			finish: cardDetails.finish,
			cardPrice: cardDetails.cardPrice,
		};

		setEditFormData(formValues);
	};

	const handleCancelClick = () => {
		setEditCardId(null);
	};

	const handleDeleteClick = (cardId) => {
		const newCard = [...cardDetails];
		const index = cardDetails.findIndex(
			(cardDetails) => cardDetails.id === cardId
		);

		newCard.splice(index, 1);
		setCardDetails(newCard);
	};

	return (
		<div className="table">
			<form onSubmit={handleEditFormSubmit}>
				<table>
					<thead>
						<tr>
							<th>Quantity</th>
							<th>Card Name</th>
							<th>Version</th>
							<th>Condition</th>
							<th>Language</th>
							<th>Finish</th>
							<th>Card Price</th>
						</tr>
					</thead>
					<tbody>
						{cardDetails.map((val) => (
							<Fragment>
								{editCardId === val.id ? (
									<EditableRow
										editFormData={editFormData}
										handleEditFormChange={handleEditFormChange}
										handleCancelClick={handleCancelClick}
									/>
								) : (
									<ReadOnlyRow
										cardDetails={val}
										handleEditClick={handleEditClick}
										handleDeleteClick={handleDeleteClick}
									/>
								)}
							</Fragment>
						))}
					</tbody>
				</table>
			</form>
			<h2>Add a new card</h2>
			<form onSubmit={handleAddFormSubmit}>
				<input
					type="text"
					name="quantity"
					required="required"
					placeholder="Please enter the quantity"
					onChange={handleAddFormChange}
				/>
				<input
					type="text"
					name="cardName"
					required="required"
					placeholder="Enter a card name"
					onChange={handleAddFormChange}
				/>
				<input
					type="text"
					name="version"
					required="required"
					placeholder="Version"
					onChange={handleAddFormChange}
				/>
				<input
					type="text"
					name="condition"
					required="required"
					placeholder="Condition"
					onChange={handleAddFormChange}
				/>
				<input
					type="text"
					name="language"
					required="required"
					placeholder="Language"
					onChange={handleAddFormChange}
				/>
				<input
					type="text"
					name="finish"
					required="required"
					placeholder="Finish"
					onChange={handleAddFormChange}
				/>
				<input
					type="text"
					name="cardPrice"
					required="required"
					placeholder="Bought card at"
					onChange={handleAddFormChange}
				/>
				<button type="submit">Add card</button>
			</form>
		</div>
	);
}

export default Table;
