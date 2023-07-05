 import React from "react";

interface palletItem {
    icon: string,
    name: string,
}

const PalletItem: React.FC<palletItem> = ({icon, name}: palletItem) => {
    return (
        <div className='pallet-item'>
            {icon}
            {name}
        </div>
    );
}

interface palletList {
    pitems: palletItem[]
}

const Pallet: React.FC<palletList> = ({pitems}) => {
    return (
     <div className="pallet">
        {pitems.map(pitem => (<PalletItem icon={pitem.icon} name={pitem.name} />))}
     </div>
    );
}

export default Pallet
