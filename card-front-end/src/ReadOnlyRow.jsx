import React from 'react';

function ReadOnlyRow({ cardDetails, handleEditClick, handleDeleteClick }) {
	return (
		<tr>
			<td>{cardDetails.quantity}</td>
			<td>{cardDetails.cardName}</td>
			<td>{cardDetails.version}</td>
			<td>{cardDetails.condition}</td>
			<td>{cardDetails.language}</td>
			<td>{cardDetails.finish}</td>
			<td>{cardDetails.cardPrice}</td>
			<td>
				<button type="button" onClick={(e) => handleEditClick(e, cardDetails)}>
					Edit
				</button>
				<button type="button" onClick={() => handleDeleteClick(cardDetails.id)}>
					Delete
				</button>
			</td>
		</tr>
	);
}

export default ReadOnlyRow;
