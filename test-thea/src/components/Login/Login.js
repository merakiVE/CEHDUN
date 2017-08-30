//Dependencies
import React, { Component } from 'react';

//Styles
import './Login.css'

//Material-UI
import Dialog from 'material-ui/Dialog';
import FlatButton from 'material-ui/FlatButton';
import TextField from 'material-ui/TextField';

const customContentStyle = {
  	width: '55%',
  	maxWidth: 'none',
};

class Modal extends Component{

	constructor(){
		super();
		this.state = {valueEmail: "",valuePasswd: "",errorEmail: "", errorPasswd: ""};

		this.handleSubmit = this.handleSubmit.bind(this);
	}

	handleSubmit = (e) => {
		let email = this.state.valueEmail;
		let passwd = this.state.valuePasswd;

		if (email === "" && passwd === ""){
			this.setState({
			  	errorEmail: "Este campo es requerido!",
			  	errorPasswd: "Este campo es requerido!",
			});
		}

		else if (email === ""){
			this.setState({
			  	errorEmail: "Este campo es requerido!",
			});
		}
		else if(passwd === ""){
			this.setState({
			  	errorPasswd: "Este campo es requerido!",
			});
		}else{
			this.setState({
			  	errorEmail: "",
			  	errorPasswd: "",
			});
		}
	};
	
	changeInputEmail = (e) => {
    	this.setState({
    		valueEmail: e.target.value,
    		errorEmail: "",
    		errorPasswd: ""
    	});
	}

	changeInputPasswd = (e) => {
    	this.setState({
    		valuePasswd: e.target.value,
    	  	errorEmail: "",
    	  	errorPasswd: ""
    	});
	}

	render(){
		return(
			<Dialog title="Iniciar Sesion" 
				contentClassName={'modal-content'} 
				contentStyle={customContentStyle} 
				modal={false} 
				open={this.props.open} 
				onRequestClose={this.props.onRequestClose} 
				autoScrollBodyContent={true}>

				<TextField
				  hintText="Correo electrónico"
				  floatingLabelText="Correo electrónico"
				  type="email"
				  value={this.state.valueEmail}
				  errorText={this.state.errorEmail}
				  onChange={this.changeInputEmail}
				/>
				<br />
				<TextField
				  hintText="Contraseña"
				  floatingLabelText="Contraseña"
				  type="password"
				  value={this.state.valuePasswd}
				  errorText={this.state.errorPasswd}
				  onChange={this.changeInputPasswd}
				/>
				<br />
				<br />
			  	<FlatButton
			  	  	label="Cancelar"
			  	  	primary={true}
			  	  	onTouchTap={this.props.onTouchTapCancel}
			  	/>

			  	<FlatButton
			  	  	label="Enviar"
			  	  	primary={true}
			  	  	onTouchTap={this.handleSubmit}
			  	/>
			</Dialog>
		);
	}
}

export default Modal;
