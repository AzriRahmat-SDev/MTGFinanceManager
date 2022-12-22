import React from 'react';

function EditableRow({
	editFormData,
	handleEditFormChange,
	handleCancelClick,
}) {
	return (
		<tr>
			<td>
				<input
					type="text"
					required="required"
					placeholder="Please enter the quantity"
					name="quantity"
					value={editFormData.quantity}
					onChange={handleEditFormChange}
				/>
			</td>
			<td>
				<input
					type="text"
					required="required"
					placeholder="Enter a card name"
					name="cardName"
					value={editFormData.cardName}
					onChange={handleEditFormChange}
				/>
			</td>
			<td>
				<input
					type="text"
					required="required"
					placeholder="Version of the card"
					name="version"
					value={editFormData.version}
					onChange={handleEditFormChange}
				/>
			</td>
			<td>
				<input
					type="text"
					required="required"
					placeholder="Condition"
					name="condition"
					value={editFormData.condition}
					onChange={handleEditFormChange}
				/>
			</td>
			<td>
				<input
					type="text"
					required="required"
					placeholder="Language"
					name="language"
					value={editFormData.language}
					onChange={handleEditFormChange}
				/>
			</td>
			<td>
				<input
					type="text"
					required="required"
					placeholder="Finish"
					name="finish"
					value={editFormData.finish}
					onChange={handleEditFormChange}
				/>
			</td>
			<td>
				<input
					type="text"
					required="required"
					placeholder="Bought card at"
					name="cardPrice"
					value={editFormData.cardPrice}
					onChange={handleEditFormChange}
				/>
			</td>
			<td>
				<button type="submit">Save</button>
				<button type="button" onClick={handleCancelClick}>
					Cancel
				</button>
			</td>
		</tr>
	);
}

export default EditableRow;
