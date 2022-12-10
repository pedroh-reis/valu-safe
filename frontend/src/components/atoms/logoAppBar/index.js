import React from 'react';
import StyledImg from './styles';
import { Images } from '../../../config';

const Logo = (props) => <StyledImg src={Images.COMPLETE_LOGO} alt="ValuSafe Logo" title="valusafe logo" {...props} />;

export default Logo;
