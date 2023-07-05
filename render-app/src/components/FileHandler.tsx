import React, { ChangeEvent, useState } from 'react';

const FileReadWriteComponent: React.FC = () => {
  const [fileContent, setFileContent] = useState('');

  const handleFileChange = (event: ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.[0];

    if (file) {
      const reader = new FileReader();

      reader.onload = (e) => {
        const content = e.target?.result as string;
        setFileContent(content);
      };

      reader.readAsText(file);
    }
  };

  const handleSaveFile = () => {
    const blob = new Blob([fileContent], { type: 'text/plain' });
    const url = URL.createObjectURL(blob);

    const link = document.createElement('a');
    link.href = url;
    link.download = 'output.txt';
    link.click();
  };

  return (
    <div>
      <h2>File Read/Write Example</h2>

      <input type="file" onChange={handleFileChange} />
      <button onClick={handleSaveFile}>Save File</button>

      <div>
        <h3>File Content:</h3>
        <pre>{fileContent}</pre>
      </div>
    </div>
  );
};

export default FileReadWriteComponent;

