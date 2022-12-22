import React from 'react';
import BoilingVerdict from './BoilingVerdict'


const scaleNames = {
    c: "Celsius",
    f: "Fahrenheit",
}

function toCelsius(fahrenheit) {
    return (fahrenheit - 32) * 5 / 9;
}

function toFahrenheit(celsius) {
    return (celsius * 9 / 5) + 32;
}

function tryConvert(temperature, convert) {
    const input = parseFloat(temperature);
    if (Number.isNaN(input)) {
        return '';
    }
    const output = convert(input);
    const rounded = Math.round(output * 1000) / 1000;
    return rounded.toString();
}

class TemperatureInput extends React.Component {

    constructor(props) {
        super(props);
        this.handleChange = this.handleChange.bind(this);
    }

    //handle every keystroke from thr user
    handleChange(e) {
        this.props.onTemperatureChange(e.target.value);
    }
    

    render() {
        const temperature = this.props.temperature
        const scale = this.props.scale
        return(
            <fieldset>
                <legend>Enter a temperature in {scaleNames[scale]}: </legend>
                <input value={temperature} onChange={this.handleChange}/>
            </fieldset>

        )
    }
}

class CalculateTemperature extends React.Component {
    constructor(props) {
        super(props);
        this.state = {temperature: '',scale: 'c'}
        this.handleChangeC = this.handleChangeC.bind(this)
        this.handleChangeF = this.handleChangeF.bind(this)
    }

    handleChangeC(temperature) {
        this.setState({scale: 'c',temperature})
    }

    handleChangeF(temperature) {
        this.setState({scale: 'f',temperature})
    }
    render(){
        const scale = this.state.scale;
        const temperature = this.state.temperature;
        const celsius = scale === 'f' ? tryConvert(temperature, toCelsius) : temperature;
        const fahrenheit = scale === 'c' ? tryConvert(temperature, toFahrenheit) : temperature;

        return (
            <div>
                <TemperatureInput scale="c" temperature={celsius} onTemperatureChange={this.handleChangeC}/>
                <TemperatureInput scale="f" temperature={fahrenheit} onTemperatureChange={this.handleChangeF}/>
                <BoilingVerdict celsius={parseFloat(celsius)}/>
            </div>
            
        )
    }
}

export default CalculateTemperature;