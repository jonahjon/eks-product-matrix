import React from 'react'
import './search.styles.css'

export const Search = ({ placeholder, handleChange }) => {
    return <input
        className='search'
        type='search'
        placeholder={placeholder}
        onChange={handleChange}
    />
};