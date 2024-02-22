import React, { useState, useEffect } from "react";
import Message from "./Message";

const Conversation = () => {
  const [data, setData] = useState(null);

  useEffect(() => {
    async function getData() {
      let res = await fetch(process.env.REACT_APP_BASE_URL);

      if (!res.ok) {
        console.log("An error Occured!");
        return;
      }
      let jsonData = await res.json();

      setData(jsonData);
    }

    getData();
  }, []);
  return (
    <div className="bg-white rounded-lg shadow">
      {data && data.map((msg, index) => <Message key={index} {...msg} />)}
    </div>
  );
};

export default Conversation;
