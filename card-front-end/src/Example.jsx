import React from 'react';
import ReactDOM from 'react-dom';

//js object declared
const user = {
	firstName: 'John',
	lastname: 'Cena',
	links: {
		youtube: 'http://youtube.com'
	}
};



const element = <h1 className="greeting">Hello, {Name(user)}</h1>;
const link = <a href="https://www.youtube.com/">link</a>

const root = ReactDOM.createRoot(document.getElementById('root'));

root.render(
	<body>
		<div>
			<h1>{element}</h1>
			<div>
			<h1>{link}</h1>
			</div>
			
		</div>
	</body>
);

function Name(val) {
	return (
		<div className="App">
			<h1>{user.firstName} {user.lastname}</h1>
		</div>
		)
}

export default Name;