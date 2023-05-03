import React, { useEffect, useState } from "react";
import Cookies from "js-cookie";

// interface Props {
//   isTrue: boolean;
//   setisTrue: React.Dispatch<React.SetStateAction<boolean>>;
// }
  
const Home = (props: {name:string, created:string}) => {
  return props.name ? (
    <p>
      Welcome: {props.name}, Account Created {props.created}
    </p>
  ) : (
    <p>Hello</p>
  );
};

export default Home;
  
 