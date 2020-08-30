import React from 'react';

import { Wrapper, Header, Body, Content } from './styles';
import { MainHeader } from '../../index'

interface Props {
  header?: JSX.Element,
  children: React.ReactNode,
}

const MainTemplate = ({ header, children }:Props) => {
  return (
    <Wrapper>
      <Header>
        {header}
      </Header>
      <Body>
        <Content>{children }</Content>
      </Body>
    </Wrapper>
  )
}

MainTemplate.defaultProps = {
  header: <MainHeader/>,
}

export default MainTemplate
