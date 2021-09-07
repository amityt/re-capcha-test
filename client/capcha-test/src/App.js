import {useState} from "react";
import ReCAPTCHA from "react-google-recaptcha";

function App() {
	const [token, setToken] = useState("")
	console.log(token);
  const handleSubmit = (e) => {
    e.preventDefault();
		const name = document.querySelector("#name").value;
		fetch("http://localhost:8000/ping", {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify({ name: name, captcha: token }),
		})
			.then((res) => {
				if (!res.ok) {
					throw res
				}
				return res.json();
			})
			.then((data) => {
				alert(`Captcha validated successfully. Welcome ${data.data}!`);
			}).catch((err)=>{
				alert(`Error validating the captcha. Try Again!`);
		});
  };

  return (
		<div>
			<form onSubmit={(e) => handleSubmit(e)}>
				<label for="name">Name</label>
				<input type="text" name="name" id="name" class="form-control" />
				<ReCAPTCHA
					sitekey="6LffWk4cAAAAAGt3aOSDIQ6mAzBUfzM6-kR0lcPh"
					onChange={(value)=>{setToken(value)}}
				/>
				<br />
				<input type="submit" value="Submit" />
			</form>
		</div>
  );
}

export default App;
